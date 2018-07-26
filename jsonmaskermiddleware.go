package main

type JSONMaskMiddleware struct {
	input   chan JSONLogObject
	outputs []chan JSONLogObject
}

func (m *JSONMaskMiddleware) Join(input chan JSONLogObject) {
	m.input = input
}

func (m *JSONMaskMiddleware) Listen() error {
	defer func() {
		// close the channel because NO MOAR STUFF
		for _, output := range m.outputs {
			close(output)
		}
	}()

	for {
		select {
		case jsonObj, ok := <-m.input:
			if !ok {
				return nil
			}

			for _, output := range m.outputs {
				jsonObj.Message = emojifier(jsonObj.Message)
				output <- jsonObj
			}
		}
	}

	return nil
}

func (m *JSONMaskMiddleware) Subscribe() chan JSONLogObject {
	ch := make(chan JSONLogObject, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}

// emojifier converts a's into a cat heart eye emoji
func emojifier(line string) string {
	find := 'a'
	replace := '😻 '
	return strings.Map(func(r rune) rune {
		if r == find {
			return replace
		}
		return r
	}, line)
}
