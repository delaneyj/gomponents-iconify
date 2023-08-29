package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infraredremote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.849 20.953L27.098 10.2c-.44-.44-1.332-.26-1.992.4L4.182 31.524c-.66.66-.839 1.552-.399 1.992l10.751 10.751c.44.44 1.332.261 1.992-.399h0L37.45 22.944c.66-.66.839-1.552.399-1.991Z"/><circle cx="14.943" cy="33.042" r="2.484" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.171 6.002a10.8 10.8 0 0 1 7.702 3.318h0a10.805 10.805 0 0 1 3.323 7.697m-10.204-7.88a7.115 7.115 0 0 1 7.068 7.077m-6.324-4.147a3.574 3.574 0 0 1 3.302 2.65a2.78 2.78 0 0 1 .089.74"/><circle cx="26.279" cy="16.303" r=".75" fill="currentColor"/><circle cx="32.222" cy="22.247" r=".75" fill="currentColor"/><circle cx="29.25" cy="19.275" r=".75" fill="currentColor"/><circle cx="23.307" cy="19.275" r=".75" fill="currentColor"/><circle cx="29.25" cy="25.218" r=".75" fill="currentColor"/><circle cx="26.279" cy="22.247" r=".75" fill="currentColor"/><circle cx="20.335" cy="22.247" r=".75" fill="currentColor"/><circle cx="26.279" cy="28.19" r=".75" fill="currentColor"/><circle cx="23.307" cy="25.218" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}