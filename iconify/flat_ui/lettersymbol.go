package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lettersymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50" cy="50" r="50" fill="#4A245A" fill-rule="evenodd" clip-rule="evenodd"/><clipPath id="flatUiLettersymbol0"><circle cx="50" cy="50" r="50"/></clipPath><g fill-rule="evenodd" clip-path="url(#flatUiLettersymbol0)" clip-rule="evenodd"><g opacity=".2"><path fill="none" stroke="#fff" stroke-width="2" d="M50 12L17 31m0 0l-33-19m33 19v38m0 0l-33 19m33-19l33 19m0 0v38m0-38l33-19m0 0l33 19M83 69V31m0 0l33-19M83 31L50 12m0 0v-45"/><circle cx="83" cy="69" r="4" fill="#fff"/><circle cx="83" cy="69" r="2" fill="#4A245A"/><circle cx="50" cy="88" r="4" fill="#fff"/><circle cx="50" cy="88" r="2" fill="#4A245A"/><circle cx="17" cy="69" r="4" fill="#fff"/><circle cx="17" cy="69" r="2" fill="#4A245A"/><circle cx="17" cy="31" r="4" fill="#fff"/><circle cx="17" cy="31" r="2" fill="#4A245A"/><circle cx="50" cy="12" r="4" fill="#fff"/><circle cx="50" cy="12" r="2" fill="#4A245A"/><circle cx="83" cy="31" r="4" fill="#fff"/><circle cx="83" cy="31" r="2" fill="#4A245A"/></g><path fill="#fff" d="M52.67 30.995c2.578 0 4.698.283 6.361.85c1.927.812 2.891 2.486 2.891 5.021c0 2.289-.729 3.884-2.188 4.782c-1.457.898-3.354 1.347-5.69 1.347H40.963v-12H52.67zm2.075 22.011c2.403.023 4.271.342 5.597.953c2.379 1.091 3.567 3.098 3.567 6.02c0 3.455-1.228 5.791-3.681 7.014c-1.353.658-3.244.99-5.67.99H40.962V53.006h13.783zm1.796-31L30 22v54.994L54.751 77c2.783 0 5.362-.247 7.741-.745c2.378-.497 4.439-1.418 6.189-2.761a15.175 15.175 0 0 0 3.879-4.327c1.628-2.588 2.44-5.51 2.44-8.769c0-3.159-.711-5.847-2.134-8.061c-1.423-2.213-3.532-3.831-6.328-4.851c1.84-.944 3.233-1.989 4.178-3.133c1.691-2.04 2.538-4.739 2.538-8.097c0-3.258-.839-6.056-2.516-8.394c-2.785-3.804-7.515-5.757-14.197-5.856z"/></g>`),
		g.Group(children),
	)
}