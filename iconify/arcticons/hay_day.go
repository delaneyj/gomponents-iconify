package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HayDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.462 25.886c.903-3.323-2.925-9.36-8.411-9.248c-4.425.092-8.24 4.15-8.059 8.455c.18 4.248 4.212 7.845 7.31 7.574c.791-.069 1.551-.546 3.039-1.497c3.78-2.417 5.67-3.626 6.12-5.284Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.497 27.473c3.056 1.341 4.311 6.262 3.254 7.98c-3.83.132-9.227-3.766-9.227-3.766"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.994 32.682c-.165 4.785 1.75 6.569 4.36 6.965s6.606.353 4.003-4.844"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.143 32.445c1.757 4.296 5.739 5.937 7.76 6.305m.684-13.7c.254-4.269 5.3-5.082 8.274-7.101c1.49-1.012 2.235-3.059 2.048-4.525c-.562-4.393-2.875-6.068-5.02-5.515c-4.228 1.09-1.916 10.139-6.21 8.124c-4.293-2.015 2.004-12.879-6.01-11.394c-4.278.793-4.789 3.39-3.567 6.11c.934 2.079 2.779 3.781 1.949 5.89m-2.147.359s-2.816-6.037-6.044-4.961c-4.095 1.365-3.678 8.81 1.145 8.995"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.076 21.865c2.954 1.269 3.725 8.138 1.044 10.417m-4.09 7.184c1.178 3.054 2.004 4.034 2.004 4.034M15.758 31.567c.99 5.207-2.048 11.933-2.048 11.933m8.091-20.697c1.076 0 1.949 1.287 1.949 2.874s-.873 2.873-1.949 2.873s-1.948-1.286-1.948-2.873H21.8v-2.874Zm7.889 6.577c1.565-1.581.054-3.113-.54-2.712c-.76.511-.472 1.565-.472 1.565"/>`),
		g.Group(children),
	)
}