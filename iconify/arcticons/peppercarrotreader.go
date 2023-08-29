package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peppercarrotreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.862 24.023c4.068-.854 14.203-5.114 19.638-9.825"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.442 22.978a5.087 5.087 0 0 1 1.406 3.767l3.466.804l.402-.854l-.955-.653l.553-.904l-.452-1.115v-1.598l-1.005-.326l-.904-.879l-.602.753l-.603-.3l-1.306 1.305a4.819 4.819 0 0 0-1.607-1.632c-.804-.327-2.436-.201-4.144-.503a13.427 13.427 0 0 1-5.65-2.787l.352 1.733l1.657 1.28l-1.783.277l-.05 1.456l1.657 1.331l-1.255.703l.125.904l2.637.779l-.879 1.03l-.05.753l1.005.327l-2.085 1.355l.176.88l2.536.15l.1 1.005l-1.657 1.255l.402.527l.879-.351a45.895 45.895 0 0 0 5.299-2.989a28.56 28.56 0 0 0 3.741-3.716"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.041 18.056l-1.356.929l.477.829l-1.532.728l-.552 1.356v2.084l.853.98l-1.431.201l.151.577l1.506.904l-.15.854l.552.929l.502.955l-.301.854l2.009 1.481l.754.804l1.424.435m1.773-8.748a18.795 18.795 0 0 0 5.567-.075m-4.872 3.943a18.156 18.156 0 0 0 3.993-1.155m3.086-2.304a6.445 6.445 0 0 0 1.434-.61"/>`),
		g.Group(children),
	)
}