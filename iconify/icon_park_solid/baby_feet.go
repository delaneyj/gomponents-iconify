package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyFeet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBabyFeet0"><g fill="#fff"><path stroke="#fff" stroke-width="4" d="M15 20.612c-1.424 6.15 6.493 7.715 4.624 12.048c-1.87 4.332-6.055 3.466-5.588 7.798c.468 4.333 6.006 4.394 11.048 1.784c10.083-5.221 12.187-16.625 7.624-21.63c-5.608-6.15-16.284-6.15-17.708 0Z"/><ellipse cx="34.535" cy="13.535" rx="2" ry="3" transform="rotate(40 34.535 13.535)"/><ellipse cx="29.381" cy="10.603" rx="2" ry="3" transform="rotate(25 29.38 10.603)"/><ellipse cx="23.381" cy="9.603" rx="2" ry="3" transform="rotate(6 23.38 9.603)"/><ellipse cx="14" cy="8" stroke="#fff" stroke-width="4" rx="3" ry="4" transform="rotate(-20 14 8)"/><ellipse cx="38.535" cy="17.536" rx="2" ry="3" transform="rotate(50 38.535 17.536)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBabyFeet0)"/>`),
		g.Group(children),
	)
}