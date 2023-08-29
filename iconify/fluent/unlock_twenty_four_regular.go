package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2.004c1.875 0 3.334 1.207 3.928 3.003a.75.75 0 0 1-1.425.47C14.102 4.263 13.185 3.505 12 3.505c-1.407 0-2.42.958-2.496 2.552l-.005.194V8h8.251a2.25 2.25 0 0 1 2.245 2.096l.005.154v9.496a2.25 2.25 0 0 1-2.096 2.245l-.154.005H6.25A2.25 2.25 0 0 1 4.005 19.9L4 19.746V10.25a2.25 2.25 0 0 1 2.096-2.245L6.25 8h1.749V6.25C8 3.712 9.71 2.004 12 2.004zM17.75 9.5H6.25a.75.75 0 0 0-.743.648l-.007.102v9.496c0 .38.282.694.648.743l.102.007h11.5a.75.75 0 0 0 .743-.648l.007-.102V10.25a.75.75 0 0 0-.648-.743L17.75 9.5zm-5.75 4a1.499 1.499 0 1 1 0 2.997a1.499 1.499 0 0 1 0-2.997z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}