package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picsay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="5.448" height="8.224" x="35.624" y="31.763" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.65" ry="2.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.655 39.986v-8.223h2.692c1.522 0 2.756 1.236 2.756 2.762s-1.234 2.761-2.756 2.761h-2.692m2.692 0l2.692 2.698m-12.353.002v-8.223h2.692c1.522 0 2.756 1.236 2.756 2.762s-1.234 2.761-2.756 2.761h-2.692"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M37.756 20.877v2.526a1.87 1.87 0 0 1-1.87 1.87h0c-.517 0-.985-.209-1.323-.547"/><path d="M37.756 17.79v3.087a1.87 1.87 0 0 1-1.87 1.871h0a1.87 1.87 0 0 1-1.871-1.87V17.79"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.392 20.69a1.87 1.87 0 0 1-1.871 1.871h0a1.87 1.87 0 0 1-1.87-1.87v-1.217a1.87 1.87 0 0 1 1.87-1.87h0a1.87 1.87 0 0 1 1.87 1.87m.001 3.087v-4.958m-9.848 4.138c.458.597 1.034.82 1.834.82h1.108a1.867 1.867 0 0 0 1.867-1.867v-.008a1.867 1.867 0 0 0-1.867-1.867h-1.222a1.869 1.869 0 0 1-1.869-1.869h0c0-1.034.839-1.873 1.873-1.873h1.102c.8 0 1.376.223 1.835.82m-5.998 5.722a1.87 1.87 0 0 1-1.624.942h0a1.87 1.87 0 0 1-1.871-1.87v-1.217a1.87 1.87 0 0 1 1.87-1.87h0c.695 0 1.3.377 1.623.938"/><circle cx="16.086" cy="15.311" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.086 17.603v4.958m-6.475 0v-7.484h2.45c1.385 0 2.508 1.126 2.508 2.514s-1.123 2.513-2.508 2.513h-2.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.253 28.893c-4.403-1.492-9.341-5.279-9.586-9.875c-.33-6.198 8.21-11.238 18.338-11.238s18.337 5.031 18.337 11.238c-.43 8.66-10.92 11.403-18.81 11.234a31.933 31.933 0 0 1-.471-.011"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.253 28.893L14.04 40.22l9.021-9.98"/>`),
		g.Group(children),
	)
}