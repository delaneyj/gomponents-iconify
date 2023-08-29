package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pocketpaint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.42 15.67L10 37.43l13.85 6.12l13.88-6.12l1.82-21.76H8.43m2.21 1.81h2"/><path fill="none" stroke="currentColor" stroke-dasharray="4.54 4.54" stroke-linecap="round" stroke-linejoin="round" d="M17.18 17.48h15.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.33 17.48h2l-.19 1.99"/><path fill="none" stroke="currentColor" stroke-dasharray="4.86 4.86" stroke-linecap="round" stroke-linejoin="round" d="m36.68 24.31l-.69 7.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.76 33.99l-.19 1.99l-1.81.84"/><path fill="none" stroke="currentColor" stroke-dasharray="2.91 2.91" stroke-linecap="round" stroke-linejoin="round" d="m31.11 38.04l-3.97 1.82"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.82 40.47l-1.82.84l-1.82-.83"/><path fill="none" stroke="currentColor" stroke-dasharray="2.93 2.93" stroke-linecap="round" stroke-linejoin="round" d="m19.52 39.25l-4-1.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.18 36.81l-1.81-.83l-.19-1.99"/><path fill="none" stroke="currentColor" stroke-dasharray="4.86 4.86" stroke-linecap="round" stroke-linejoin="round" d="m11.73 29.15l-.68-7.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.83 19.47l-.19-1.99m4.67-1.81l2.93-11.12l9.12 2.37l-2.22 8.75m-7.63-8.33l4 1.09m-5.03 2.82l3.95 1.09m-4.43.85l2.68.71M17 9.33l2.58.69m-4.11 5.06l2.08.59m9.78 0L28.4 13l2.33 1l-.38 1.73m-1.57-2.63l.71-1.91l1.57.55l-.73 2m-.84-2.55c-.12-1.81.11-5.41 3.07-6.32c1.35 2.45.2 5.08-1.5 6.87m4.6 2.06l-.93 1.87m-2.64 0l1.35-3m-.28-.34l3.43-3.41l-.45 4.82a3 3 0 0 1-3-1.41"/>`),
		g.Group(children),
	)
}