package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungBrowserLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.25 37.196c-7.333 2.934-16.213-.761-18.784-8.94m-.19-7.846C12.3 12.74 20.22 8.222 27.852 10.381c7.633 2.16 12.011 10.158 9.718 17.752"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.84 30.5c-2.423.033-5.288-.333-8.277-1.134c-7.566-2.028-13.056-6.074-12.263-9.038"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.604 19.785c4.23 2.93 6.541 6.402 5.738 9.399c-.237.884-.73 1.67-1.438 2.347m-12.882 2.93a39.965 39.965 0 0 1-7.481-1.28C10.858 30.319 3.3 23.887 4.659 18.817c.809-3.021 4.6-4.878 9.79-5.28"/><circle cx="35.977" cy="35.255" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.977 31.255v8h4"/>`),
		g.Group(children),
	)
}