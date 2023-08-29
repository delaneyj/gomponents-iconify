package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InnerShadowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInnerShadowRight0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M24 44a19.938 19.938 0 0 0 14.142-5.858A19.938 19.938 0 0 0 44 24a19.938 19.938 0 0 0-5.858-14.142A19.937 19.937 0 0 0 24 4A19.938 19.938 0 0 0 9.858 9.858A19.938 19.938 0 0 0 4 24a19.937 19.937 0 0 0 5.858 14.142A19.938 19.938 0 0 0 24 44Z" clip-rule="evenodd"/><path stroke="#000" d="M33.9 33.9A13.957 13.957 0 0 0 38 24a13.96 13.96 0 0 0-4.1-9.9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInnerShadowRight0)"/>`),
		g.Group(children),
	)
}