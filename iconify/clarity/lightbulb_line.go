package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2.25a11 11 0 0 0-11 11a10.68 10.68 0 0 0 1 4.63a16.36 16.36 0 0 0 1.12 1.78a17 17 0 0 1 2 3.47a16.19 16.19 0 0 1 .59 4h2a18.17 18.17 0 0 0-.71-4.69a18.46 18.46 0 0 0-2.22-3.92a15.79 15.79 0 0 1-1-1.54A8.64 8.64 0 0 1 9 13.23a9 9 0 0 1 18.07 0a8.64 8.64 0 0 1-.86 3.77a15.79 15.79 0 0 1-1 1.54a18.46 18.46 0 0 0-2.21 3.9a18.17 18.17 0 0 0-.71 4.71h2a16.19 16.19 0 0 1 .59-4a17 17 0 0 1 2-3.47A16.31 16.31 0 0 0 28 17.86a10.68 10.68 0 0 0 1-4.63A11 11 0 0 0 18 2.25Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18.63 15.51a.8.8 0 0 0-1.13 0l-3 3l2.86 3.13v5.54H19V21l-2.24-2.45l1.89-1.89a.8.8 0 0 0-.02-1.15Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M23.86 29.15H12.11a.8.8 0 1 0 0 1.6h11.75a.8.8 0 0 0 0-1.6Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M22 32.15h-8a.8.8 0 1 0 0 1.6h8a.8.8 0 1 0 0-1.6Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="m17.32 10.89l-2.73 2.73a.8.8 0 0 0 1.13 1.13L18.45 12a.8.8 0 1 0-1.13-1.13Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}