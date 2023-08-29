package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kindle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.944 37.03a56.3 56.3 0 0 1 9.696-.751c4.318 0 11.836 1.626 20.316 4.879"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.167 36.547c1.52-.212 3.833-2.679 3.833-2.679a15.199 15.199 0 0 0 2.237-1.57l2.179 2.124v1.24l1.29.885l3.589-2.96s-.379-1.293-1.262-1.64c-.042-.62-2.748-5.06-3-5.425m-6.589.533a14.834 14.834 0 0 0 2.44 1.542c.392.028 6.532-4.093 6.532-4.093a4.73 4.73 0 0 0 2.13-1.122a2.733 2.733 0 0 0 .225-1.15s.364-.476.364-.645s.28-1.01.28-1.093"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.855 24.504s.701.869.925.869s5.103-2.58 5.215-3.14m-10.828 6.532c.09-.084 2.644-1.444 2.644-1.444"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.154 21.868c.477 1.907 3.673 5.453 3.673 5.453l3.365-3.294s-2.832-3.028-3.169-3.953s-.785-2.524-1.682-2.916s-4.458-2.243-4.458-2.243l-.673 1.233a11.73 11.73 0 0 0-5.13 9.673c0 6.561.756 8.804.756 8.804s4.85-1.037 6.589-2.888s2.742-2.972 2.742-2.972l.353-3.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.597 34.44a3.435 3.435 0 0 0 1.423 1.93m19.701-2.31a15.47 15.47 0 0 0 3.391 4.532a17.14 17.14 0 0 0 2.25 1.304M30.03 36.28c.573.756 2.396 2.817 2.956 3.448M20.167 17.079l.454-.426s2.047 1.402 2.664 1.01s.589-.842.589-.842s1.01-.617.981-.925a1.743 1.743 0 0 1 0-.449h.841s-.532-1.598-.196-2.186l.336-.59l.401.365a2.6 2.6 0 0 0 1.03-2.58c-.393-1.57-2.776-4.037-3.87-4.205s-.953.196-.953.196s-1.99-1.177-3.224-.084s-2.692 2.356-2.663 4.598s.056 2.888.224 3.113a2.61 2.61 0 0 0 .766.481l-.494.964m17.924.738l3.028.2l-5.284 5.663l-2.511-.308l4.767-5.555z"/>`),
		g.Group(children),
	)
}