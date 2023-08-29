package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCity0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42h40"/><rect width="8" height="16" x="8" y="26" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="square" stroke-linejoin="round" stroke-width="4" d="M12 34h1"/><rect width="24" height="38" x="16" y="4" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="2"/><path fill="#000" d="M22 10h4v4h-4zm8 0h4v4h-4zm-8 7h4v4h-4zm8 0h4v4h-4zm0 7h4v4h-4zm0 7h4v4h-4z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCity0)"/>`),
		g.Group(children),
	)
}