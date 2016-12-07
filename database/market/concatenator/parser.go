package concatenator

import (
	"errors"
	"strconv"
)

type parser struct {
	result            *Region
	data, word        []byte
	nonStringWord     bool
	inWord            bool
	objectCount       uint8
	skipToVal         bool
	skipToNextOrder   bool
	key               string
	err               error
	debugBreak        bool
	ii                int
	price             float64
	volume, stationId uint64
	itemId            uint64
}

func newParser(r *Region) *parser {
	this := new(parser)
	this.result = r
	return this
}

func (this *parser) parse(b []byte) error {
	this.data = b
	return this.iterator()
}

func (this *parser) iterator() error {
ITERATOR:
	for ii, ch := range this.data {
		this.ii = ii

		switch {
		case this.err != nil:
			break ITERATOR
		case ii == 0 && ch != '[':
			this.err = errors.New("JSON for market data is not an array")
		case ii == 0:
			continue
		case this.objectCount > 1:
			this.err = errors.New("Market data does not contain nested objects")
		case this.skipToNextOrder && ch != '}':
			continue
		case ch == '"' || this.inWord && this.nonStringWord && (ch == ',' || ch == '}'):
			this.wordToggle(ch)
		case this.inWord:
			this.addChar(ch)
		case ch == 't' || ch == 'T' || ch == 'f' || ch == 'F' || ch >= '0' && ch <= '9':
			this.wordToggle(ch)
		case ch == '{' || ch == '[':
			this.objectCount++
		case ch == '}':
			this.checkObjComplete(ch)
		case ch == ',':
			this.handleComma()
		case ch == ':' && this.skipToVal:
			continue
		case this.skipToVal:
			continue
		case ch == ']' && ii == (len(this.data)-1):
			{
				return nil
			}
		default:
			this.err = errors.New("Unhandled character '" + string(ch) + "'")
		}
	}

	if this.err != nil {
		return this.err
	}
	return nil
}

func (this *parser) addChar(ch byte) {
	this.skipToVal = false
	if ch >= 'A' && ch <= 'Z' {
		//dirty way to convert everything to LC
		ch += 'a' - 'A'
	}
	this.word = append(this.word, ch)
}

func (this *parser) handleComma() {
	if this.inWord {
		this.addChar(',')
		return
	}
}

func (this *parser) wordToggle(ch byte) {
	this.skipToVal = false
	if this.inWord {
		this.nonStringWord = false
		this.wordComplete(ch)
		return
	}

	this.inWord = true
	switch {
	case ch == 'f' || ch == 'F':
		fallthrough
	case ch == 't' || ch == 'T':
		fallthrough
	case ch >= '0' && ch <= '9':
		{
			this.addChar(ch)
			this.nonStringWord = true
		}
	case ch == '"':
		this.nonStringWord = false
	default:
		this.err = errors.New("Unknown word terminator '" + string(ch) + "'")
	}
}

func (this *parser) wordComplete(ch byte) {
	word := string(this.word)
	key := this.key
	this.word = nil
	this.inWord = false
	this.key = ""

	if key == "" {
		if ch == '}' {
			this.err = errors.New("unexpected end of object at index" + string(this.ii))
			return
		}
		this.key = word
		this.skipToVal = true
		return
	}

	switch key {
	case "buy":
		{
			isBuy, err := strconv.ParseBool(word)
			if err != nil {
				this.err = err
				return
			}
			if isBuy {
				this.skipOrder()
				return
			}
			this.checkObjComplete(ch)
		}
	case "issued":
		{
			this.skipToVal = true
			this.checkObjComplete(ch)
		}
	case "volume":
		{
			this.volume, this.err = strconv.ParseUint(word, 10, 64)
			this.checkObjComplete(ch)
		}
	case "price":
		{
			this.price, this.err = strconv.ParseFloat(word, 64)
			this.checkObjComplete(ch)
		}
	case "stationid":
		{
			this.stationId, this.err = strconv.ParseUint(word, 10, 64)
			this.checkObjComplete(ch)
		}
	case "type":
		{
			this.itemId, this.err = strconv.ParseUint(word, 10, 64)
			this.checkObjComplete(ch)
		}
	case "duration", "id", "minvolume", "volumeentered", "range":
		this.checkObjComplete(ch)
	default:
		this.err = errors.New("unknown pair '" + key + "' '" + word + "'")
	}

}

func (this *parser) checkObjComplete(ch byte) {
	if ch != '}' {
		//Object not finished yet
		return
	}

	if this.skipToNextOrder {
		this.skipToNextOrder = false
		return
	}

	if this.itemId == 0 || this.volume == 0 || this.price == 0 || this.stationId == 0 {
		this.err = errors.New("Missing values for object ending at " + string(this.ii))
		return
	}

	this.result.Add(this.price, this.volume, this.stationId, this.itemId)
	this.itemId = 0
	this.volume = 0
	this.stationId = 0
	this.itemId = 0
	this.objectCount--
}

func (this *parser) skipOrder() {
	this.skipToNextOrder = true
	this.itemId = 0
	this.volume = 0
	this.stationId = 0
	this.itemId = 0
	this.objectCount--
}
