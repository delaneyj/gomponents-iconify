package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandageAdhesiveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M26.828 8.444a9 9 0 0 1 12.728 12.728L21.172 39.557A9 9 0 0 1 8.444 26.829L26.828 8.444Zm11.314 1.414a7 7 0 0 0-9.9 0L9.859 28.243a7 7 0 1 0 9.9 9.9l18.384-18.385a7 7 0 0 0 0-9.9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M24 16.929a2 2 0 0 1 2.828 0l4.243 4.242a2 2 0 0 1 0 2.829L24 31.07a2 2 0 0 1-2.828 0l-4.243-4.242a2 2 0 0 1 0-2.828l7.07-7.071Zm5.657 5.657l-4.243-4.243l-7.071 7.071l4.243 4.243l7.07-7.071Z" clip-rule="evenodd"/><path d="M33.121 16.121a1 1 0 1 1-1.414-1.413a1 1 0 0 1 1.414 1.413Zm1-4a1 1 0 1 1-1.414-1.413a1 1 0 0 1 1.414 1.413Zm-4 1a1 1 0 1 1-1.414-1.413a1 1 0 0 1 1.414 1.413Zm6 6a1 1 0 1 1-1.414-1.413a1 1 0 0 1 1.414 1.413Zm1-4a1 1 0 1 1-1.414-1.413a1 1 0 0 1 1.414 1.413ZM14.707 31.707a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415Zm-1 4a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415Zm4-1a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415Zm-6-6a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415Zm-1 4a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415Z"/></g>`),
		g.Group(children),
	)
}