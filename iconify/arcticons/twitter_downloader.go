package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterDownloader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.358 18.77v.703c0 7.076-5.369 15.185-15.193 15.185c-2.9.007-5.738-.836-8.165-2.424c.421.057.847.083 1.272.077c2.404.007 4.74-.795 6.633-2.277a5.313 5.313 0 0 1-4.989-3.717c.333.07.672.106 1.012.106c.475 0 .948-.064 1.405-.19a5.32 5.32 0 0 1-4.279-5.235v-.07c.74.413 1.57.64 2.417.66a5.299 5.299 0 0 1-1.68-7.07a15.164 15.164 0 0 0 11.02 5.58a4.483 4.483 0 0 1-.148-1.222a5.305 5.305 0 0 1 9.254-3.732a10.955 10.955 0 0 0 3.394-1.3a5.375 5.375 0 0 1-2.382 3A11.154 11.154 0 0 0 37 15.96a10.934 10.934 0 0 1-2.641 2.811Z"/><circle cx="38.5" cy="38.499" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.432 43.326A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 3.383-.781 6.583-2.173 9.43M38.5 42.499v-8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.203 39.203l3.297 3.296l3.297-3.296"/>`),
		g.Group(children),
	)
}