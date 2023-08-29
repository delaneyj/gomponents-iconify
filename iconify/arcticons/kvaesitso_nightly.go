package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KvaesitsoNightly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.5" cy="38.499" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 42.499v-8l5.3 8v-8M28.983 20.862h5.638V26.5h-5.638zm-7.936 0h5.638V26.5h-5.638zm-7.937 0h5.638V26.5H13.11zm-4.474-5.103l13.46-6.715a4.152 4.152 0 0 1 3.66-.023L39.365 15.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.926 19.778v7.807c0 .958.777 1.734 1.735 1.734h24.287c.958 0 1.734-.776 1.734-1.734v-7.807c0-.958-.776-1.735-1.735-1.735H11.661c-.958 0-1.735.777-1.735 1.735ZM34.639 32.66a1.99 1.99 0 0 0-1.14-.357H14.11a2 2 0 0 0-2 2v7.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.432 43.326A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 3.383-.781 6.583-2.173 9.43"/>`),
		g.Group(children),
	)
}