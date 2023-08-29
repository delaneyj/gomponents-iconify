package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M11.09 14.57h.31a6.43 6.43 0 0 1 .09-2a2.09 2.09 0 1 1 1.47-3a6.58 6.58 0 0 1 1.55-1.31a4.09 4.09 0 1 0-3.42 6.33Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M13 18.14a6.53 6.53 0 0 1-1.28-2.2h-.63a8.67 8.67 0 0 0-6.43 2.52l-.24.28v7h2v-6.23a7 7 0 0 1 4.67-1.6a8.09 8.09 0 0 1 1.91.23Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M31.35 18.42A8.59 8.59 0 0 0 25 15.91c-.32 0-.6 0-.9.06a6.53 6.53 0 0 1-1.35 2.25a7.9 7.9 0 0 1 2.25-.31a6.94 6.94 0 0 1 4.64 1.58v6.27h2V18.7Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="M18.1 19.73a9.69 9.69 0 0 0-7.1 2.74l-.25.28v7.33a1.57 1.57 0 0 0 1.61 1.54h11.47a1.57 1.57 0 0 0 1.61-1.54v-7.35l-.25-.28a9.58 9.58 0 0 0-7.09-2.72Zm5.33 9.88h-10.7v-6.06a8.08 8.08 0 0 1 5.37-1.82a8 8 0 0 1 5.33 1.8Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted"/><path fill="currentColor" d="M20.28 14.27a2.46 2.46 0 1 1-2.42-2.89a2.44 2.44 0 0 1 1 .24a3.67 3.67 0 0 1 .43-2a4.41 4.41 0 0 0-1.48-.27A4.47 4.47 0 1 0 22.14 15a3.69 3.69 0 0 1-1.86-.73Z" class="clr-i-outline--alerted clr-i-outline-path-5--alerted"/><path fill="currentColor" d="m27.18.8l-5.72 9.91a1.28 1.28 0 0 0 1.1 1.91H34a1.28 1.28 0 0 0 1.1-1.91L29.39.8a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-6--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}