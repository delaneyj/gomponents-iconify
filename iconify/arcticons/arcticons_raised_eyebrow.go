package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArcticonsRaisedEyebrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24 10.84l-11.4 6.58v13.16L24 37.16l11.4-6.58V17.42L24 10.84Z"/><circle cx="8.07" cy="14.8" r="3.11"/><circle cx="24" cy="5.61" r="3.11"/><circle cx="39.93" cy="14.8" r="3.11"/><circle cx="39.93" cy="33.2" r="3.11"/><circle cx="24" cy="42.39" r="3.11"/><circle cx="8.07" cy="33.2" r="3.11"/><path d="m12.6 30.58l-1.84 1.06M24 37.16v2.12m0-28.44V8.71m11.4 21.87l1.84 1.06M35.4 17.42l1.84-1.06M12.6 17.42l-1.84-1.06"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="20.106" cy="23.016" r="2.198"/><path d="M30.631 22.62c0-1.432-.68-2.593-2.194-2.593c-1.548 0-2.195 1.16-2.195 2.593s.647 2.852 2.195 2.852c1.585 0 2.194-1.42 2.194-2.852Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 28.628h11m-11.855-9.616c1.072-.313 1.79-.405 2.755-.327c.967.078 1.647.39 2.565 1.337m8.676-1.362c-1.295-1.557-2.36-2.3-3.125-2.3c-1.12 0-1.577.35-2.14.902"/>`),
		g.Group(children),
	)
}