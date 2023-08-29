package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZhihuishuTreenity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.426 16.053C20.289 27.2 14.502 38.32 9.988 40.22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.595 26.13c7.9-3.12 12.786-3.855 18.187-4.997m-8.688 5.604c3.82 2.205 5.076 6.43 7.54 9.797M14.903 7.78a9.411 9.411 0 0 1-2.253 6.758c3.088-2.941 7.232-3.348 11.387-3.728m15.245 8.055l.123 10.814m-8.056-9.734l3.748-.062m-2.54 1.926l-1.7 2.007m2.458-.369l-.123 4.977m1.987-.676c-1.347.736-3.723.855-4.035 2.212m2.117-4.342l2.308-.027m-4.66.768l2.334.002m1.611-4.875a1.681 1.681 0 0 1 .9 1.721m1.537-1.721l.159 3.602m-5.929-2.056l1.794.415m4.638 5.454c.339.879.832 1.1 1.453.778"/>`),
		g.Group(children),
	)
}