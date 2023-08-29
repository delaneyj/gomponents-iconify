package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InnerShadowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInnerShadowUp0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" d="M24 44a19.938 19.938 0 0 0 14.142-5.858A19.938 19.938 0 0 0 44 24a19.938 19.938 0 0 0-5.858-14.142A19.937 19.937 0 0 0 24 4A19.938 19.938 0 0 0 9.858 9.858A19.938 19.938 0 0 0 4 24a19.937 19.937 0 0 0 5.858 14.142A19.938 19.938 0 0 0 24 44Z" clip-rule="evenodd"/><path d="M33.9 14.1A13.956 13.956 0 0 0 24 10a13.956 13.956 0 0 0-9.9 4.1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInnerShadowUp0)"/>`),
		g.Group(children),
	)
}