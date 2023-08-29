package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpenseRegister(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m16.517 14.344l7.705-4.8l10.274 8.688v12.566l-5.967 4.836V23.816l-12.012-9.472Zm9.541-5.086L31.9 5.646l10.46 7.293l-6.433 4.926m.277 10.748l6.296-5.14m-6.296 2.479l6.296-5.14m-6.296 2.48l6.296-5.14m-6.296 2.48l6.296-5.14"/><path d="m35.314 14.172l2.723-2.077l-1.865-1.247l-1.498 1.131M5.5 31.954l13.543 10.4l7.423-5.91"/><path d="m5.5 29.285l13.543 10.4l7.423-5.91"/><path d="m5.604 26.616l13.543 10.401l7.423-5.91"/><path d="m5.59 23.948l13.542 10.4l7.423-5.91m-6.32-4.688c-.226 1.027-1.694 1.554-3.278 1.175h0c-1.584-.378-2.685-1.517-2.459-2.545c.226-1.027 1.694-1.553 3.278-1.175s2.685 1.518 2.459 2.545Z"/><path d="m15.051 15.826l-9.295 5.595l13.331 10.117l7.64-6.015"/></g>`),
		g.Group(children),
	)
}