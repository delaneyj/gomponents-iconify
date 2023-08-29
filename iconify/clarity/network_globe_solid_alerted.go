package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkGlobeSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M26.58 32h-18a1 1 0 1 0 0 2h18a1 1 0 0 0 0-2Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M31.73 15.4h-6.17a18.87 18.87 0 0 1-1.62 2.52a2.33 2.33 0 0 1 .33 1.19a22 22 0 0 0 5 .45a11.88 11.88 0 0 1-.61 1.53h-.56a17.41 17.41 0 0 1-4.32-.56a2.29 2.29 0 0 1-3 .62a18.43 18.43 0 0 1-7 3.5a2.34 2.34 0 0 1-1.57 1.79l-.29.06a11.93 11.93 0 0 1-3.39-2.8h.66a2.33 2.33 0 0 1 4.37-.58A16.94 16.94 0 0 0 19.78 20a2.32 2.32 0 0 1-.18-1.17c-.42-.24-.84-.49-1.25-.76a17.53 17.53 0 0 1-5.35-5.6a2.31 2.31 0 0 1-2.28-.63a27.31 27.31 0 0 0-5 4.74v-.57a12 12 0 0 1 .14-1.73a18.75 18.75 0 0 1 4.2-3.8a2.28 2.28 0 0 1 1.1-2.25c-.12-.43-.24-.86-.33-1.3c0-.14 0-.29-.11-.64a12 12 0 0 1 1.37-.87c.1.59.14.9.21 1.21s.2.85.32 1.27h.25a2.33 2.33 0 0 1 1.13.63a18.59 18.59 0 0 1 6.39-1L23 3A14 14 0 0 0 3.75 16c0 .45 0 .89.07 1.33A14 14 0 0 0 31.76 16c0-.2-.02-.4-.03-.6Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted"/><path fill="currentColor" d="M14.26 11.64a16 16 0 0 0 4.93 5.23c.34.23.69.43 1 .63a2.28 2.28 0 0 1 2.58-.57a17.29 17.29 0 0 0 1-1.54h-1.6A3.68 3.68 0 0 1 19 9.89l.56-.89a17.08 17.08 0 0 0-4.84.88a2.25 2.25 0 0 1-.47 1.77Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}