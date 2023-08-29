package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phimpme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.717" cy="21.231" r="5.308" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.563 19.063l-6.105-.012m-1.259 2.18l3.049-5.281m-4.844 2.18l3.047 5.268m-3.58 0l6.1.013m-1.8 3.098l3.062-5.273m1.787 3.105l-3.044-5.29m2.393 8.579a9.15 9.15 0 0 0 4.132-9.882a11.225 11.225 0 0 0-4.629-6.46a19.482 19.482 0 0 0-7.51-2.948a31.069 31.069 0 0 0-10.601-.2h0a39.073 39.073 0 0 1 11.402-3.348A23.581 23.581 0 0 1 33.79 6.124a14.712 14.712 0 0 1 5.167 3.738a11.798 11.798 0 0 1 2.767 5.715A10.661 10.661 0 0 1 28.369 27.63Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.786 43.5l-2.913-3.496l-.328 2.538a30.236 30.236 0 0 1-5.399-16.16a21.824 21.824 0 0 1 .542-5.969a14.112 14.112 0 0 1 2.554-5.392a12.06 12.06 0 0 1 10.612-4.51a12.335 12.335 0 0 1 8.373 4.534a6.941 6.941 0 0 0-9.421 1.519a8.477 8.477 0 0 0-1.53 3.591a14.864 14.864 0 0 0-.144 3.926a37.488 37.488 0 0 0 .985 5.695q.763 3.206 1.768 6.347l1.166 4.077l-2.777-1.646a3.992 3.992 0 0 1 .942 3.017a3.946 3.946 0 0 1-.307 1.148l-3.866-3.458l1.57 3.606l-3.04-.924"/>`),
		g.Group(children),
	)
}