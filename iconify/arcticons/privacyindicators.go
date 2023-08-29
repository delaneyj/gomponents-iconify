package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Privacyindicators(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="15.901" height="34.846" x="16.049" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.167"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.41 39.345V43.5h3.148v-4.155"/><circle cx="24" cy="32.23" r="3.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="21.923" r="3.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="11.616" r="3.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.049 9.149h-4.494s.855 2.442 4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.049 14.962H13.9v-3.72m2.149 7.774h-4.494s.855 2.442 4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.049 24.829H13.9v-3.72m2.149 7.775h-4.494s.855 2.442 4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.049 34.697H13.9v-3.72M31.951 9.149h4.494s-.855 2.442-4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.951 14.962H34.1v-3.72m-2.149 7.774h4.494s-.855 2.442-4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.951 24.829H34.1v-3.72m-2.149 7.775h4.494s-.855 2.442-4.494 2.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.951 34.697H34.1v-3.72"/>`),
		g.Group(children),
	)
}