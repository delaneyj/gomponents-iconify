package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Concern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConcern0"><g fill="none" stroke="#fff"><path stroke-linecap="round" stroke-width="4" d="M10.858 9.858A19.937 19.937 0 0 0 5 24a19.937 19.937 0 0 0 5.858 14.142m28.284 0A19.937 19.937 0 0 0 45 24a19.937 19.937 0 0 0-5.858-14.142M34.9 33.9A13.956 13.956 0 0 0 39 24a13.96 13.96 0 0 0-4.1-9.9m-19.8 0A13.956 13.956 0 0 0 11 24a13.96 13.96 0 0 0 4.1 9.9"/><path fill="#fff" stroke-linejoin="round" stroke-width="3.5" d="M28.182 20C30.29 20 32 21.612 32 23.6c0 2.588-2.546 4.8-3.818 6c-.849.8-1.91 1.6-3.182 2.4c-1.273-.8-2.333-1.6-3.182-2.4c-1.273-1.2-3.818-3.412-3.818-6c0-1.988 1.71-3.6 3.818-3.6c1.328 0 2.498.64 3.182 1.61c.684-.97 1.854-1.61 3.182-1.61Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConcern0)"/>`),
		g.Group(children),
	)
}