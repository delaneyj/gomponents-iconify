package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jefit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.194" cy="8.094" r="3.594" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.824 37.617c-.116-.044-.667-.169-1.19-.334a14.013 14.013 0 0 1-1.114-.413c-6.276-2.62-9.73-7.79-9.73-14.71q0-.836.081-1.668m33.899-2.197a17.207 17.207 0 0 1 .44 3.865h0a17.21 17.21 0 0 1-8.02 14.552M9.8 12.438a17.21 17.21 0 0 1 18.394-6.969"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.238 32.5c2.452 7.433 6.94 11 11.51 11a6.562 6.562 0 0 0 2.943-.466a11.934 11.934 0 0 0 1.405-.68a11.174 11.174 0 0 0 1.864-1.512a11.817 11.817 0 0 0 1.63-1.939a38.074 38.074 0 0 0 4.544-10.198c.782-2.782 1.353-5.624 1.972-8.447c.181-.836.33-1.662.43-2.384a16.06 16.06 0 0 0 3.438-1.762a15.44 15.44 0 0 0-2.41-4.576c-4.348 2.336-10.63 3.675-17.232 3.675a44.085 44.085 0 0 1-11.337-1.493a15.651 15.651 0 0 0-1.801 4.789a42.892 42.892 0 0 0 12.984 1.87a60.54 60.54 0 0 0 9.24-.64c-1.212 11.557-5.656 19.705-10.79 19.705c-2.524 0-5.276-1.848-7.248-5.495"/>`),
		g.Group(children),
	)
}