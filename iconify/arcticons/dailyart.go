package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailyart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="13.557" height="19.545" x="14.99" y="9.468" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.779"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.547 22.235v20.336M21.09 29.013v3.728c-4.8 0-12.088 6.27-12.088 9.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 42.57s-2.937-8.981-7.642-8.981V16.21A10.782 10.782 0 0 0 24.077 5.43h0a10.781 10.781 0 0 0-10.782 10.781v17.378C9.737 34.888 6.008 39.972 5.5 42.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.029 22.178c.263.874 2.523.933 2.895 0m-2.895-4.848c-.283-.874-2.712-.933-3.112 0m9.263 0c-.283-.874-2.712-.933-3.112 0m-3.915 7.136a3.164 3.164 0 0 0 2.324.763a3.178 3.178 0 0 0 2.407-.89"/><circle cx="18.075" cy="17.68" r=".75" fill="currentColor"/><circle cx="24.43" cy="17.68" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}