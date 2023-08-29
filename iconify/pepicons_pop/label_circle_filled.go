package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopLabelCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M7.91 7.564c-.018.283-.034.596-.047.927c-.054 1.358-.047 2.907.102 4.078l7.265 7.265l4.95-4.95l-7.265-7.265c-1.17-.149-2.72-.156-4.078-.102c-.331.013-.643.03-.927.047ZM6.094 5.748c-.097.092-.49 4.567-.077 7.33c.044.294.189.56.399.77l7.753 7.754a1.5 1.5 0 0 0 2.122 0l5.657-5.657a1.5 1.5 0 0 0 0-2.122L14.194 6.07a1.382 1.382 0 0 0-.77-.4c-2.763-.411-7.238-.02-7.33.078Z" clip-rule="evenodd"/><path d="M12.142 10.466a1 1 0 1 1-1.415 1.414a1 1 0 0 1 1.415-1.414Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopLabelCircleFilled0)"/></g>`),
		g.Group(children),
	)
}