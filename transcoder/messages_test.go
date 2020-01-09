package transcoder

import "testing"

func TestSerializeTranscodeCommand(t *testing.T) {
	command := TranscodeCommand{
		FileName: "foo",
		GuildId:  "guild",
	}
	want := `{"fileName":"foo","guildId":"guild"}`
	assertSerializes(command, want, t)
}

func TestDeserializeTranscodeCommand(t *testing.T) {
	data := []byte(`{"fileName":"foo","guildId":"guild"}`)
	got, err := ToTranscodeCommand(data)
	if err != nil {
		t.Errorf("Error while deserializing: %v", err)
	}
	want := TranscodeCommand{
		FileName: "foo",
		GuildId:  "guild",
	}
	if got.FileName != want.FileName {
		t.Errorf("Incorrect %v value. Got %v, expected %v", "FileName", got.FileName, want.FileName)
	}
	if got.GuildId != want.GuildId {
		t.Errorf("Incorrect %v value. Got %v, expected %v", "GuildId", got.GuildId, want.GuildId)
	}
}

func assertSerializes(message Message, want string, t *testing.T) {
	got, err := message.Serialize()

	if err != nil {
		t.Errorf("Error while serializing: %v", err)
	}

	if string(got) != want {
		t.Errorf("Serialize() = %v; want %v", string(got), want)
	}
}
