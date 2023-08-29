package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m754 511l-272 66q-18-4-27.5-18t-4.5-29l155-185q6-17 30-22.5t39 15.5l80 111q26 39 0 62zM487 641l247 127q19 0 29.5 19.5T757 826L633 951q-20 16-39 5.5T574 927L448 681q-1-17 11-28.5t28-11.5zm-170 378l-168-67q-24-13-21-36.5t21-29.5l168-166q22-23 44.5-11t22.5 44v232q0 25-21 35.5t-46-1.5zm0-521L85 140q-18-6-21-29.5T85 74L317 7q25-12 46-1.5T384 41v424q0 32-22.5 44T317 498zm-13 74q16 15 16 37t-16 37L76 757q-11 11-29 12t-32.5-10.5T0 725V493q0-22 14.5-33.5T47 449t29 12z"/>`),
		g.Group(children),
	)
}