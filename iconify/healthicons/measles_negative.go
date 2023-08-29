package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeaslesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsMeaslesNegative0)" clip-rule="evenodd"><path d="M40 24c0 8.837-7.163 16-16 16c-.089 0-.178 0-.266-.002a2 2 0 1 0-3.667-.485a16.035 16.035 0 0 1-10.994-9.74a2 2 0 1 0-.795-2.79A16.087 16.087 0 0 1 8 24a15.96 15.96 0 0 1 5.32-11.914a2 2 0 1 0 3.285-2.278A15.93 15.93 0 0 1 24 8c3.107 0 6.007.885 8.461 2.418a1.5 1.5 0 1 0 2.359 1.795A15.958 15.958 0 0 1 40 24Zm-25.779-7.628a1 1 0 0 1 1.309-.22l4 2.5a1 1 0 0 1 0 1.696l-4 2.5a1 1 0 0 1-1.219-1.573c.982-.934 1.174-1.422 1.171-1.729c-.002-.33-.225-.85-1.198-1.848a1 1 0 0 1-.063-1.326Zm18.249-.22a1 1 0 0 1 1.246 1.546c-.973.998-1.196 1.518-1.198 1.848c-.003.307.189.795 1.171 1.73a1 1 0 0 1-1.219 1.572l-4-2.5a1 1 0 0 1 0-1.696l4-2.5ZM33.258 31c-1.72-3.562-5.22-6-9.258-6c-4.038 0-7.538 2.438-9.258 6c-.469.97.316 2 1.394 2h15.728c1.078 0 1.863-1.03 1.394-2ZM24.5 14a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm12 15a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/><path d="M48 0H0v48h48V0Zm-6 24c0 9.941-8.059 18-18 18S6 33.941 6 24S14.059 6 24 6s18 8.059 18 18Z"/></g><defs><clipPath id="healthiconsMeaslesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}