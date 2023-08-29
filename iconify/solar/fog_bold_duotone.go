package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FogBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 21.25a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5H8Z"/><path d="M12.476 2C9.32 2 6.762 4.528 6.762 7.647c0 .69.125 1.35.354 1.962a4.356 4.356 0 0 0-.83-.08C3.919 9.53 2 11.426 2 13.765c0 .522.096 1.023.271 1.485h18.92A5.57 5.57 0 0 0 22 12.353c0-2.472-1.607-4.573-3.845-5.338C17.837 4.194 15.415 2 12.476 2Z" opacity=".5"/><path d="M2 15.249a.75.75 0 0 0 0 1.5h20a.75.75 0 0 0 0-1.5H2Z"/><path d="M5 18.25a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5H5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}