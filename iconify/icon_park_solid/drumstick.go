package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drumstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDrumstick0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m14.15 33.82l-1.413 9.9l-8.486-8.486l9.9-1.414Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24.05 6.95l-9.9 9.9c-4.686 4.686-4.686 12.284 0 16.97v0c4.687 4.687 12.285 4.687 16.97 0l9.9-9.9"/><ellipse cx="32.535" cy="15.435" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="12" ry="7" transform="rotate(45 32.535 15.435)"/><circle cx="30.061" cy="11.398" r="2" fill="#000" transform="rotate(45 30.06 11.398)"/><circle cx="37.132" cy="18.469" r="2" fill="#000" transform="rotate(45 37.132 18.47)"/><circle cx="31.475" cy="17.055" r="2" fill="#000" transform="rotate(45 31.475 17.055)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDrumstick0)"/>`),
		g.Group(children),
	)
}