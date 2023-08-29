package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airlift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" d="M14.865 22.813v-1.54a8.036 8.036 0 0 1 8.013-8.013h0a8.039 8.039 0 0 1 7.32 4.763M8.108 34.69h0A3.619 3.619 0 0 1 4.5 31.081v-2.103a3.619 3.619 0 0 1 3.608-3.608h0a3.619 3.619 0 0 1 3.607 3.608v2.103a3.618 3.618 0 0 1-3.607 3.607Zm3.607-9.319v9.318m3.282-9.318v9.318m15.458-10.457v10.457m-12.132-9.318v9.318m0-6.279a3.619 3.619 0 0 1 3.608-3.607h0a3.647 3.647 0 0 1 .723.073m5.087 9.809c-.044.002-.087.004-.132.004h-.214a2.502 2.502 0 0 1-2.495-2.494v-12.1"/><circle cx="30.455" cy="20.784" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.18 34.74V22.742a2.545 2.545 0 0 1 2.545-2.544h0a4.962 4.962 0 0 1 .823.063m2.589 1.846v10.815a1.817 1.817 0 0 0 1.818 1.817h.545m-10.451-9.633h9.997"/>`),
		g.Group(children),
	)
}