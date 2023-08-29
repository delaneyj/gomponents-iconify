package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrecioLuz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.624 32.435c-.258-.879-.211-1.918-.43-2.938a9.623 9.623 0 0 0-1.642-3.859c-6.032-8.084.221-14.931 6.223-15.022c8.634.044 11.29 7.659 8.624 13.023a37.903 37.903 0 0 0-2.286 4.034a13.232 13.232 0 0 0-1.134 4.815a28.556 28.556 0 0 1-9.358-.057"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.624 32.435a1.802 1.802 0 0 0-.117 2.669a75.626 75.626 0 0 0 9.623.194a1.998 1.998 0 0 0-.147-2.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.812 35.325a1.786 1.786 0 0 0 .137 2.086a42.956 42.956 0 0 0 8.671.144a1.8 1.8 0 0 0 .134-2.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.418 37.682a1.432 1.432 0 0 0 .154 1.811c.738.704 4.027.734 5.217.09c.658-.526.6-1.08.37-1.823m1.1 2.561h6.068m-2.37-12.557h11.467m-9.331-3.746h6.917c2.434-.057 2.558 1.422 2.568 1.831v7.404m-5.217-11.424v3.353m1.123 10.703l.953 2.652l2.404-.643M6.513 17.118l2.053.416m-.889-6.459l2.274 1.633m2.444-6.997l1.844 3.638m6.04-.889L20.26 4.5m8.034 1.257l-2.19 3.488m4.269 3.588l2.639-1.838m-1.083 6.418l2.364-.325"/><circle cx="28.272" cy="30.815" r=".75" fill="currentColor"/><circle cx="32.249" cy="30.802" r=".75" fill="currentColor"/><circle cx="36.239" cy="30.802" r=".75" fill="currentColor"/><circle cx="28.248" cy="36.039" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.613 23.167c.647-.017 1.217.885 1.663.885c.637-.034 1.318-.899 1.968-.858c.564-.014 1.328.885 1.858.858c.59 0 .872-.902 1.61-.939"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.276 24.052l-2.045 7.93l-2.092-7.944"/><circle cx="36.618" cy="38.631" r="4.869" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}