package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mercuryssh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.089 24.68l-1.995.22c-1.108 1.467-5.963 7.112-8.927 8.832a4.34 4.34 0 0 1-2.249.657l-2.84-.09c-.956-.03-2.18 2.14-.388 2.843l5.215 2.048c1.923.755 4.256.258 6.177-.503c1.774-.703 2.567-3.555 4.476-3.57l5.213-.04c2.195-.016 3.18-3.472 2.354-5.18c-1.03-2.131-4.573-2.186-6.892-1.713c-1.006.206-2.46 1.855-2.46 1.855"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.068 28.802c3.208 1.022 4.608 2.741 4.49 6.316M19.094 24.9C12.304 10.855 26.97 17.673 38.4 7.291c1.18 1.882.493 4.418-.368 5.959"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.305 27.726c-6.79-14.044 8.815-6.423 19.608-17.532c1.287-.147 1.125 4.263-.445 5.674c1.44-.203 1.26 3.07-1.447 5.013c1.878.638-.68 3.25-2.588 3.708c.805.849-.23 2.36-3.753 3.375c.29 2.173-2.492 2.041-5.461 1.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.468 15.868a29.762 29.762 0 0 1-10.97 3.139m9.522 1.874a26.193 26.193 0 0 1-9.971.586m7.383 3.123a26.336 26.336 0 0 1-8.495-.955m4.743 4.329a16.388 16.388 0 0 1-6.362-2.567"/>`),
		g.Group(children),
	)
}