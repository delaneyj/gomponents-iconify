package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="m11.806 14.129l-3.484 3.484c-1.16 1.161-2.322 0-1.16-1.162l9.29-9.29c1.16-1.161 2.322 0 1.16 1.161l-3.483 3.484s2.322 0 4.645 2.323l2.323 2.322l9.037-9.037a2 2 0 0 1 2.828 0l1.151 1.15a2 2 0 0 1 .453 2.133l-5.34 13.884l6.968 9.29l4.03-2.015a.954.954 0 0 1 1.101 1.529l-7.94 7.94a.954.954 0 0 1-1.53-1.102l2.016-4.03l-9.29-6.967l-13.884 5.34a2 2 0 0 1-2.132-.453l-1.15-1.15a2 2 0 0 1 0-2.829l9.036-9.037l-2.322-2.323c-2.323-2.323-2.323-4.645-2.323-4.645Z"/>`),
		g.Group(children),
	)
}