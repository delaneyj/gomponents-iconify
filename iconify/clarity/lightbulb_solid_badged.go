package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.86 29.15H12.11a.8.8 0 1 0 0 1.6h11.75a.8.8 0 0 0 0-1.6Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M22 32.15h-8a.8.8 0 1 0 0 1.6h8a.8.8 0 1 0 0-1.6Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M22.5 6a7.47 7.47 0 0 1 .5-2.56a11 11 0 0 0-16 9.8a10.68 10.68 0 0 0 1 4.63a16.36 16.36 0 0 0 1.12 1.78a17 17 0 0 1 2 3.47a16.19 16.19 0 0 1 .59 4h5.69v-5.51l-2.86-3.13l3-3a.8.8 0 0 1 1.13 1.13l-1.89 1.89L19 21v6.17h5.3a16.19 16.19 0 0 1 .59-4a17 17 0 0 1 2-3.47A16.31 16.31 0 0 0 28 17.86a10.63 10.63 0 0 0 1-4.43A7.5 7.5 0 0 1 22.5 6Zm-4 6l-2.73 2.73a.8.8 0 1 1-1.13-1.13l2.73-2.73A.8.8 0 1 1 18.45 12Z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}