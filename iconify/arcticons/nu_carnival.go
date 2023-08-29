package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NuCarnival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.523 8.748c-3.364 3.364-8.916 15.056-6.224 23.551c-4.963-2.86-8.51-13.81-4.71-20.103C4.336 15.297 4.85 22.878 5.5 24.561c-1.621-.673-2.778-3.65-2.778-3.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.56 20.523C19.122 19.43 30.375 8.748 30.674 3.558M2.946 28.374c1.633 2.187 2.867 3.897 5.362 3.533c1.066 4.486 9.998 12.562 14.21 13.542m22.917-23.132c-.696 3.87-6.958 10.907-9.93 11.524c3.869-4.542 6.842-14.07 6.66-21.348"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.924 22.194c.485-4.273-1.354-11.764-2.635-14.259a43.977 43.977 0 0 1-8.868 22.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.272 22.569c-2.382 4.372-7.655 8.356-12.366 10.487c2.664-2.551 7.121-13.346 6.112-23.916m-6.833 28.942a11.84 11.84 0 0 0 6.392 2.986a6.84 6.84 0 0 0 4.416-2.124m-1.956 2.86a4.178 4.178 0 0 1-1.948.82m2.861-14.161a20.66 20.66 0 0 0 1.359 6.675a13.515 13.515 0 0 0-1.359 1.774m-1.887-.491l.358.189M7.425 27.294s.393 2.86.519 3.14c-.407-.462-1.654-1.57-1.654-1.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.262 26.467a6.066 6.066 0 0 0-3.491-3.196a8.528 8.528 0 0 0-6.14.87a4.944 4.944 0 0 0 4.893 3.728m14.676-.14c1.445 0 6.781-2.103 6.781-6.233a9.185 9.185 0 0 0-7.471 2.635m-15.053-5.936c2.608.113 5.996 1.5 7.207 2.749m14.908-4.914a27.332 27.332 0 0 0-4.99 3.432m-10.675 3.55c1.373-.554 6.42-3.162 7.037-8.937"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.162 23.017c-1.973 0-2.435 4.004.151 4.004a2.314 2.314 0 0 0 2.092-3.512m11.511 1.741c-.173.825 1.72 2.114 2.953.6s.898-3.63-.175-3.322m-.913 2.201l.093-.14m-15.291.308l.093-.14m19.037 7.153c-.657 2.8-7.82 12.315-8.486 12.674"/>`),
		g.Group(children),
	)
}