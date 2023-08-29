package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.91 7.564c-.018.283-.034.596-.047.927c-.054 1.358-.047 2.907.102 4.078l7.265 7.265l4.95-4.95l-7.265-7.265c-1.17-.149-2.72-.156-4.078-.102c-.331.013-.643.03-.927.047ZM6.094 5.748c-.097.092-.49 4.567-.077 7.33c.044.294.189.56.399.77l7.753 7.754a1.5 1.5 0 0 0 2.122 0l5.657-5.657a1.5 1.5 0 0 0 0-2.122L14.194 6.07a1.382 1.382 0 0 0-.77-.4c-2.763-.411-7.238-.02-7.33.078Z" clip-rule="evenodd"/><path d="M12.142 10.466a1 1 0 1 1-1.415 1.414a1 1 0 0 1 1.415-1.414Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}