package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.285 38.506h-1.17s-.18.9.09 1.979c.36 1.08-.72.18-.72.18c-1.169-1.98-1.619-3.239-1.978-5.398c-1.62-9.805 4.947-20.24 4.587-21.5s-1.17-.99-1.17-.99c-4.407 1.08-6.296 13.224-6.746 16.913c0 .27-.45.36-.54.09c-.9-1.98-1.17-4.138-1.35-6.208c-.989-9.715 6.208-18.351 16.553-19.88c2.34-.36 3.599-.18 5.218.27c.45.09.54.629.27.899c-11.695 11.695-12.325 29.147-12.325 32.925c0 .45-.36.72-.72.72Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.415 5.04c.36-.18.27-.809-.09-.809c-7.827-1.26-22.67 22.58-16.283 39.942c.36.45.9.45.72-.09c-.54-2.249.54-2.789 1.62-2.879c.629-.09 1.169-.72 1.079-1.349c-.9-19.341 9.265-28.517 12.324-30.856c.45-.36.45-.99 0-1.26l-.9-.72c-.81-.629.63-1.439 1.53-1.978Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.463 11.158c-.45.54-.45 1.26-.09 1.89a18.62 18.62 0 0 1 2.429 9.175c0 8.366-5.488 15.473-13.134 17.902c-.36.09-.54.36-.72.63c-.36.81-.18 1.619.27 2.788c-7.107-6.747 2.699-33.645 12.954-33.645c-.63.09-1.35.81-1.71 1.26Z"/>`),
		g.Group(children),
	)
}