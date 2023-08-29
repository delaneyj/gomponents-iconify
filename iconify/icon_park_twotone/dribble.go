package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dribble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDribble0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 24a19.938 19.938 0 0 1-5.889 14.173A19.937 19.937 0 0 1 24 44C12.954 44 4 35.046 4 24a19.932 19.932 0 0 1 5.5-13.775A19.943 19.943 0 0 1 24 4a19.937 19.937 0 0 1 14.111 5.827A19.938 19.938 0 0 1 44 24Z"/><path d="M44 24c-2.918 0-10.968-1.1-18.173 2.063C18 29.5 12.333 34.831 9.863 38.147"/><path d="M16.5 5.454C19.63 8.343 26.46 15.698 29 23c2.54 7.302 3.48 16.28 4.06 18.835"/><path d="M4.154 21.5c3.778.228 13.779.433 20.179-2.3c6.4-2.733 11.907-7.76 13.796-9.355M5.5 31.613a20.076 20.076 0 0 0 9 9.991"/><path d="M4 24a19.932 19.932 0 0 1 5.5-13.775M24 4a19.943 19.943 0 0 0-14.5 6.225M32 5.664a20.037 20.037 0 0 1 6.111 4.163A19.938 19.938 0 0 1 44 24c0 2.462-.445 4.821-1.26 7M24 44a19.937 19.937 0 0 0 14.111-5.827"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDribble0)"/>`),
		g.Group(children),
	)
}