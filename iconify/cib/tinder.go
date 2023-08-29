package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.421 12.599a.335.335 0 0 0 .496.089l.02-.016c.521-.437.984-.88 1.427-1.36C18.64 6.672 16.145.969 16.125.917a.663.663 0 0 1 .197-.792a.66.66 0 0 1 .823.041c14.5 13.484 10.657 23.755 10.381 24.448c-1.163 4.188-5.781 7.219-11.261 7.375c-.183.011-.323.011-.484.011c-6.469 0-11.969-3.984-11.969-9.079v-.077c0-7.063 6.391-14.032 6.677-14.323a.584.584 0 0 1 .704-.141a.626.626 0 0 1 .385.599c-.057 1.381.224 2.584.843 3.6z"/>`),
		g.Group(children),
	)
}