package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkGlobeSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M26.58 32h-18a1 1 0 1 0 0 2h18a1 1 0 0 0 0-2Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M31.5 13.35a7.54 7.54 0 0 1-1.5.15a7.46 7.46 0 0 1-3.28-.76a18.59 18.59 0 0 1-2.78 5.18a2.33 2.33 0 0 1 .33 1.19a22 22 0 0 0 5 .45a11.88 11.88 0 0 1-.61 1.53h-.56a17.41 17.41 0 0 1-4.32-.56a2.29 2.29 0 0 1-3 .62a18.43 18.43 0 0 1-7 3.5a2.34 2.34 0 0 1-1.57 1.79l-.29.06a11.93 11.93 0 0 1-3.39-2.8h.66a2.33 2.33 0 0 1 4.37-.58A16.94 16.94 0 0 0 19.78 20a2.32 2.32 0 0 1-.18-1.17c-.42-.24-.84-.49-1.25-.76a17.53 17.53 0 0 1-5.35-5.6a2.31 2.31 0 0 1-2.28-.63a27.31 27.31 0 0 0-5 4.74v-.57a12 12 0 0 1 .14-1.73a18.75 18.75 0 0 1 4.2-3.8a2.28 2.28 0 0 1 1.1-2.25c-.12-.43-.24-.86-.33-1.3c0-.14 0-.29-.11-.64a12 12 0 0 1 1.37-.87c.1.59.14.9.21 1.21s.2.85.32 1.27h.25a2.33 2.33 0 0 1 1.13.63a18.56 18.56 0 0 1 8.65-.87a7.45 7.45 0 0 1 .41-4.59A14 14 0 0 0 3.75 16c0 .45 0 .89.07 1.33a14 14 0 1 0 27.68-4Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M14.72 9.87a2.25 2.25 0 0 1-.47 1.77a16 16 0 0 0 4.93 5.23c.34.23.69.43 1 .63a2.28 2.28 0 0 1 2.58-.57a17 17 0 0 0 2.61-5a7.52 7.52 0 0 1-2.16-2.67a17 17 0 0 0-8.49.61Z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}