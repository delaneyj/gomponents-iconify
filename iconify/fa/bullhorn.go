package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 512q53 0 90.5 37.5T1792 640t-37.5 90.5T1664 768v384q0 52-38 90t-90 38q-417-347-812-380q-58 19-91 66t-31 100.5t40 92.5q-20 33-23 65.5t6 58t33.5 55t48 50T768 1438q-29 58-111.5 83T488 1532.5T356 1477q-7-23-29.5-87.5t-32-94.5t-23-89t-15-101t3.5-98.5T282 896H160q-66 0-113-47T0 736V544q0-66 47-113t113-47h480q435 0 896-384q52 0 90 38t38 90v384zm-128 604V162q-394 302-768 343v270q377 42 768 341z"/>`),
		g.Group(children),
	)
}