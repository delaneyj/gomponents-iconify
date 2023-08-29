package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mybl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.183 33.068c-.755 3.774-1.593 5.227-2.348 6.68m-6.931.336c3.661-5.199 8.217-12.773 8.469-20.096c-.615-1.118-.615-6.932.419-10.482c-2.096 6.932-3.634 9.587-4.724 11.572s-3.628 6.89-5.757 10.034c-2.292 3.382-3.522 5.394-4.808 8.636m-2.263-3.633c2.291-2.6 4.723-9.196 5.98-11.879s3.187-9.475 8.86-15.988m-3.968.783c-2.683 2.208-3.929 6.704-4.696 8.19c-3.913 7.574-10.565 18.782-14.59 22.527M29.66 7.969c-.629 2.795-3.843 8.748-3.843 8.748m1.369-7.686c-3.074 2.6-8.664 12.494-10.62 19.565S8.515 40 8.542 40.084c5.171-7.49 4.137-11.6 4.305-16.854S15.25 7.969 15.25 7.969c-3.13 1.9-7.155 6.904-7.155 6.904"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.304 16.214c-3.298 3.578-2.32 12.578-2.32 12.578M7.849 9.758C8.04 8.36 9.076 7.8 9.076 7.8m11.207.169c-3.103 3.522-4.696 14.534-4.696 14.534"/>`),
		g.Group(children),
	)
}