package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 2h-.066c-.886 0-1.65 0-2.262.082c-.655.088-1.284.287-1.793.797c-.51.51-.709 1.138-.797 1.793C3 5.284 3 6.048 3 6.934V7.95a1 1 0 1 0 2 0V7c0-.971.002-1.599.064-2.061c.059-.434.153-.57.229-.646c.076-.076.212-.17.646-.229C6.4 4.002 7.029 4 8 4h8c.971 0 1.599.002 2.061.064c.434.059.57.153.646.229c.076.076.17.212.229.646C18.998 5.4 19 6.029 19 7v.95a1 1 0 1 0 2 0V6.934c0-.886 0-1.65-.082-2.262c-.088-.655-.287-1.284-.797-1.793c-.51-.51-1.138-.709-1.793-.797C17.716 2 16.952 2 16.066 2H8Z" clip-rule="evenodd"/><path d="M13 4h-2v16h2V4Z" opacity=".5"/><path fill-rule="evenodd" d="M6 21a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}