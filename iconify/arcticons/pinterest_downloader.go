package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinterestDownloader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.5" cy="38.499" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.432 43.326A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 3.383-.781 6.583-2.173 9.43M38.5 42.499v-8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.203 39.203l3.297 3.296l3.297-3.296M19.356 34.5s3.522-14.666 3.87-15.586c-4.135.31-2.71 8.2 1.525 9.068c3.226.755 6.8-2.983 6.748-7.13s-2.49-6.216-5.572-7.077c-3.082-.86-7.984.204-9.285 4.515a6.919 6.919 0 0 0 1.971 7.536"/>`),
		g.Group(children),
	)
}