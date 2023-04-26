package botstuff

import "testing"

func TestQr(t *testing.T) {
	str, err := generateQrCodeB64EncodedForUrl("bla")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if str != goldenTestResult {
		t.Error("Wanted something different")
	}
}

const goldenTestResult = `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgAAAAIAAQMAAADOtka5AAAABlBMVEX///8AAABVwtN+AAABk0lEQVR42uzbQXKqMBzAYRwXLD2CR/FoejSP4jFcvDFvCMQQqrFVLJ3x+61oab/lf5IAjSRJkiRJkj6qVSg7NU2z7i4O8fZ6cvsIAAAAbwUO+ad9ArrLEM795TUAAAB4COzitLmMgNFEavJtAAAAWAIoLgEAAGBRoPvtNi2yAAAA4PtAsWvbhH/95bPbPgAAAHgemBxMR2A00n58sg0AAABPAnfrR9oLAQCAjwOGNVI3hlZpIoW0MGr74XQEAACA3wXiQVEzWNfW2aq/DQQAAACvA3E5dYx/eulXViHv2kK22nsn2wAAADDZnw2l4XR993loG87d7Kps8AAAAGAWYLQw2qf/mjy8P91/VAYAAABzA1UrpENuAAAAqAG3vjIdJlJxMN3m42oAAAB4H/Dl/cSQPlMdjpQejTQAAACYEah+FNYPuvqzNgAAALgJbNIYKg6mAwAAACwAFM/HYtWDaQAAAJgXKHZtxTlSSI/K2vT6IgAAAFSA21+ZhvxR2C7dBAAAgLcCkiRJkiRJ+sP9DwAA//+cMiMctqAECgAAAABJRU5ErkJg`
