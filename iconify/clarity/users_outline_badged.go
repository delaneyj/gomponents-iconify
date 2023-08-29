package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M11.09 14.57h.31a6.43 6.43 0 0 1 .09-2a2.09 2.09 0 1 1 1.47-3a6.58 6.58 0 0 1 1.55-1.31a4.09 4.09 0 1 0-3.42 6.33Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M13 18.14a6.53 6.53 0 0 1-1.28-2.2h-.63a8.67 8.67 0 0 0-6.43 2.52l-.24.28v7h2v-6.23a7 7 0 0 1 4.67-1.6a8.09 8.09 0 0 1 1.91.23Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M31.35 18.42A8.59 8.59 0 0 0 25 15.91c-.32 0-.6 0-.9.06a6.53 6.53 0 0 1-1.35 2.25a7.9 7.9 0 0 1 2.25-.31a6.94 6.94 0 0 1 4.64 1.58v6.27h2V18.7Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M17.86 18.3a4.47 4.47 0 1 0-4.47-4.47a4.47 4.47 0 0 0 4.47 4.47Zm0-6.93a2.47 2.47 0 1 1-2.47 2.47a2.47 2.47 0 0 1 2.47-2.47Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M18.1 19.73a9.69 9.69 0 0 0-7.1 2.74l-.25.28v7.33a1.57 1.57 0 0 0 1.61 1.54h11.47a1.57 1.57 0 0 0 1.61-1.54v-7.35l-.25-.28a9.58 9.58 0 0 0-7.09-2.72Zm5.33 9.88h-10.7v-6.06a8.08 8.08 0 0 1 5.37-1.82a8 8 0 0 1 5.33 1.8Z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><path fill="currentColor" d="M26.37 12a2 2 0 0 1-2.09.42a6.53 6.53 0 0 1 .15 1.38a6.59 6.59 0 0 1 0 .68a4 4 0 0 0 .57.06a4.08 4.08 0 0 0 3.3-1.7a7.45 7.45 0 0 1-1.93-.84Z" class="clr-i-outline--badged clr-i-outline-path-6--badged"/><path fill="currentColor" d="M22.95 6.93a4.16 4.16 0 0 0-1.47 1.44A6.59 6.59 0 0 1 23 9.77a2.1 2.1 0 0 1 .59-.83a7.44 7.44 0 0 1-.64-2.01Z" class="clr-i-outline--badged clr-i-outline-path-7--badged"/><circle cx="30.33" cy="5.67" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-8--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}