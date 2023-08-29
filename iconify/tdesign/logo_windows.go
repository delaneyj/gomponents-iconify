package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 1.847V12H12V3.276l10-1.429ZM14 5.01V10h6V4.153l-6 .857Zm-3-1.188V12H2V5.303l9-1.48ZM4 7.001v3h5V6.177L4 7Zm-2 6h9V21.1l-9-.857V13Zm2 2v3.423l5 .476V15H4Zm8-2h10v9.156l-10-1.459V13Zm2 2v3.968l6 .875V15h-6Z"/>`),
		g.Group(children),
	)
}