package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beetv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.58 12.35a7.12 7.12 0 0 1 10.07 0m-10.07 0c-.619 2.177.576 3.932 3.357 6.713"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.58 12.35L18.695 8.1A8.728 8.728 0 0 0 7.72 9.947a8.86 8.86 0 0 0-.42 11.215c2.587 3.306 7.019 4.02 10.89 2.399l10.746-4.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.629 38.371c-1.294-1.374-3.482-7.206 1.298-14.544"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.605 36.527s-1.123 3.99-1.86 4.728M8.11 31.004l4.443 4.443M9.929 25.52l6.275 6.276m3.271-3.271l-4.195-4.196m8.108.283l-2.274-2.274"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M35.216 11.948s3.442-3.373 2.444-5.406s-3.418-.554-3.917-.035"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.65 22.42a7.12 7.12 0 0 0 0-10.07m0 10.07c-2.177.619-3.932-.576-6.713-3.357"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.65 22.42l4.25 6.885a8.728 8.728 0 0 1-1.849 10.974a8.86 8.86 0 0 1-11.215.42c-3.305-2.586-4.02-7.019-2.398-10.89l4.5-10.746"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.629 38.371c1.374 1.294 7.206 3.482 14.544-1.298"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.473 39.395s-3.99 1.123-4.728 1.86m10.251-1.365l-4.443-4.443m9.927 2.624l-6.276-6.275m3.271-3.271l4.196 4.195m-.283-8.108l2.274 2.274"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M36.052 12.784s3.373-3.442 5.406-2.445s.554 3.42.035 3.918"/>`),
		g.Group(children),
	)
}