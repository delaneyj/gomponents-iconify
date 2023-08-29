package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheKanaQuiz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.189 5.565c-.621 6.508-1.031 12.669.685 15.313"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.982 8.271a43.994 43.994 0 0 0 12.743-.788m12.71-1.918c-.826 5.19-1.38 10.29-4.933 15.655M38.971 8.237A28.108 28.108 0 0 1 42.5 14.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.714 10.566a65.692 65.692 0 0 1 8.496-.445c1.053.048 2.403.967 2.432 2.021a22.114 22.114 0 0 1-.925 7.331a2.006 2.006 0 0 1-2.021 1.508l-2.5-.206m6.542 6.372v8.461c.055 1.537-.623 1.405-1.199 1.405h-2.021m7.365-5.653l-16.717 1.268"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.722 27.626s.054 10.316.034 11.682a2.343 2.343 0 0 0 2.57 2.432l7.433-.205M20.096 25.228l1.713 2.467m-3.837-1.747l1.713 2.57m-6.475-1.885l4.556 10.243l-3.802-.411"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.947 30.333c5.087.267 9.752.4 13.292-.788M7.386 36.5c-2.199 8.053 6.36 5.645 9.592 5.446m-.548-31.414c-.78 4.77-4.541 7.708-6.646 9.112c-1.275.851-4.167 1.873-4.282-1.096c-.076-1.958 2.05-4.195 3.802-5.481c2.357-1.73 6.149-2.295 8.736-.754c3.79 2.258 2.47 5.598.308 7.468a9.91 9.91 0 0 1-3.87 2.022"/>`),
		g.Group(children),
	)
}