package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.5 8.5V8h-.39l-.095.379l.485.121Zm2 0l.485-.121L5.89 8H5.5v.5Zm.69 2.758l.484-.122l-.485.122Zm-3.38 0l.486.12l-.485-.12ZM9.5 10.5l.277-.416A.5.5 0 0 0 9 10.5h.5Zm3 2l-.277.416A.5.5 0 0 0 13 12.5h-.5Zm-3-8H9V5h.5v-.5ZM1.5 1h12V0h-12v1Zm12.5.5v12h1v-12h-1ZM13.5 14h-12v1h12v-1ZM1 13.5v-12H0v12h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0ZM3 3h3V2H3v1Zm0 2h3V4H3v1Zm0 2h3V6H3v1Zm.5 2h2V8h-2v1Zm1.515-.379l.69 2.758l.97-.243l-.69-2.757l-.97.242ZM5.219 12H3.781v1h1.438v-1Zm-1.923-.621l.69-2.758l-.971-.242l-.69 2.757l.97.243ZM3.78 12a.5.5 0 0 1-.485-.621l-.97-.243A1.5 1.5 0 0 0 3.78 13v-1Zm1.923-.621A.5.5 0 0 1 5.22 12v1a1.5 1.5 0 0 0 1.455-1.864l-.97.243ZM10 13v-2.5H9V13h1Zm-.777-2.084l3 2l.554-.832l-3-2l-.554.832ZM13 12.5V10h-1v2.5h1ZM9 6v3h1V6H9Zm3 0v3h1V6h-1ZM9.5 8h3V7h-3v1Zm.5-3.5v-1H9v1h1Zm3-.5h-1.5v1H13V4Zm-1.5 0h-2v1h2V4Zm-.5-.5v1h1v-1h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 2v1Zm-.5.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 9 3.5h1Z"/>`),
		g.Group(children),
	)
}