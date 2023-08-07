// Code generated by gensnowflake. DO NOT EDIT.

package discord

import "time"

// AppID is the snowflake type for a AppID.
type AppID Snowflake

// NullAppID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullAppID = AppID(NullSnowflake)

func (s AppID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *AppID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s AppID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s AppID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s AppID) IsNull() bool { return Snowflake(s).IsNull() }

func (s AppID) Time() time.Time   { return Snowflake(s).Time() }
func (s AppID) Worker() uint8     { return Snowflake(s).Worker() }
func (s AppID) PID() uint8        { return Snowflake(s).PID() }
func (s AppID) Increment() uint16 { return Snowflake(s).Increment() }

// AttachmentID is the snowflake type for a AttachmentID.
type AttachmentID Snowflake

// NullAttachmentID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullAttachmentID = AttachmentID(NullSnowflake)

func (s AttachmentID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *AttachmentID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s AttachmentID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s AttachmentID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s AttachmentID) IsNull() bool { return Snowflake(s).IsNull() }

func (s AttachmentID) Time() time.Time   { return Snowflake(s).Time() }
func (s AttachmentID) Worker() uint8     { return Snowflake(s).Worker() }
func (s AttachmentID) PID() uint8        { return Snowflake(s).PID() }
func (s AttachmentID) Increment() uint16 { return Snowflake(s).Increment() }

// AuditLogEntryID is the snowflake type for a AuditLogEntryID.
type AuditLogEntryID Snowflake

// NullAuditLogEntryID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullAuditLogEntryID = AuditLogEntryID(NullSnowflake)

func (s AuditLogEntryID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *AuditLogEntryID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s AuditLogEntryID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s AuditLogEntryID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s AuditLogEntryID) IsNull() bool { return Snowflake(s).IsNull() }

func (s AuditLogEntryID) Time() time.Time   { return Snowflake(s).Time() }
func (s AuditLogEntryID) Worker() uint8     { return Snowflake(s).Worker() }
func (s AuditLogEntryID) PID() uint8        { return Snowflake(s).PID() }
func (s AuditLogEntryID) Increment() uint16 { return Snowflake(s).Increment() }

// ChannelID is the snowflake type for a ChannelID.
type ChannelID Snowflake

// NullChannelID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullChannelID = ChannelID(NullSnowflake)

func (s ChannelID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *ChannelID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s ChannelID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s ChannelID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s ChannelID) IsNull() bool { return Snowflake(s).IsNull() }

func (s ChannelID) Time() time.Time   { return Snowflake(s).Time() }
func (s ChannelID) Worker() uint8     { return Snowflake(s).Worker() }
func (s ChannelID) PID() uint8        { return Snowflake(s).PID() }
func (s ChannelID) Increment() uint16 { return Snowflake(s).Increment() }

// CommandID is the snowflake type for a CommandID.
type CommandID Snowflake

// NullCommandID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullCommandID = CommandID(NullSnowflake)

func (s CommandID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *CommandID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s CommandID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s CommandID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s CommandID) IsNull() bool { return Snowflake(s).IsNull() }

func (s CommandID) Time() time.Time   { return Snowflake(s).Time() }
func (s CommandID) Worker() uint8     { return Snowflake(s).Worker() }
func (s CommandID) PID() uint8        { return Snowflake(s).PID() }
func (s CommandID) Increment() uint16 { return Snowflake(s).Increment() }

// EmojiID is the snowflake type for a EmojiID.
type EmojiID Snowflake

// NullEmojiID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullEmojiID = EmojiID(NullSnowflake)

func (s EmojiID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *EmojiID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s EmojiID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s EmojiID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s EmojiID) IsNull() bool { return Snowflake(s).IsNull() }

func (s EmojiID) Time() time.Time   { return Snowflake(s).Time() }
func (s EmojiID) Worker() uint8     { return Snowflake(s).Worker() }
func (s EmojiID) PID() uint8        { return Snowflake(s).PID() }
func (s EmojiID) Increment() uint16 { return Snowflake(s).Increment() }

// GuildID is the snowflake type for a GuildID.
type GuildID Snowflake

// NullGuildID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullGuildID = GuildID(NullSnowflake)

func (s GuildID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *GuildID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s GuildID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s GuildID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s GuildID) IsNull() bool { return Snowflake(s).IsNull() }

func (s GuildID) Time() time.Time   { return Snowflake(s).Time() }
func (s GuildID) Worker() uint8     { return Snowflake(s).Worker() }
func (s GuildID) PID() uint8        { return Snowflake(s).PID() }
func (s GuildID) Increment() uint16 { return Snowflake(s).Increment() }

// IntegrationID is the snowflake type for a IntegrationID.
type IntegrationID Snowflake

// NullIntegrationID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullIntegrationID = IntegrationID(NullSnowflake)

func (s IntegrationID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *IntegrationID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s IntegrationID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s IntegrationID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s IntegrationID) IsNull() bool { return Snowflake(s).IsNull() }

func (s IntegrationID) Time() time.Time   { return Snowflake(s).Time() }
func (s IntegrationID) Worker() uint8     { return Snowflake(s).Worker() }
func (s IntegrationID) PID() uint8        { return Snowflake(s).PID() }
func (s IntegrationID) Increment() uint16 { return Snowflake(s).Increment() }

// InteractionID is the snowflake type for a InteractionID.
type InteractionID Snowflake

// NullInteractionID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullInteractionID = InteractionID(NullSnowflake)

func (s InteractionID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *InteractionID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s InteractionID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s InteractionID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s InteractionID) IsNull() bool { return Snowflake(s).IsNull() }

func (s InteractionID) Time() time.Time   { return Snowflake(s).Time() }
func (s InteractionID) Worker() uint8     { return Snowflake(s).Worker() }
func (s InteractionID) PID() uint8        { return Snowflake(s).PID() }
func (s InteractionID) Increment() uint16 { return Snowflake(s).Increment() }

// MessageID is the snowflake type for a MessageID.
type MessageID Snowflake

// NullMessageID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullMessageID = MessageID(NullSnowflake)

func (s MessageID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *MessageID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s MessageID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s MessageID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s MessageID) IsNull() bool { return Snowflake(s).IsNull() }

func (s MessageID) Time() time.Time   { return Snowflake(s).Time() }
func (s MessageID) Worker() uint8     { return Snowflake(s).Worker() }
func (s MessageID) PID() uint8        { return Snowflake(s).PID() }
func (s MessageID) Increment() uint16 { return Snowflake(s).Increment() }

// RoleID is the snowflake type for a RoleID.
type RoleID Snowflake

// NullRoleID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullRoleID = RoleID(NullSnowflake)

func (s RoleID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *RoleID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s RoleID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s RoleID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s RoleID) IsNull() bool { return Snowflake(s).IsNull() }

func (s RoleID) Time() time.Time   { return Snowflake(s).Time() }
func (s RoleID) Worker() uint8     { return Snowflake(s).Worker() }
func (s RoleID) PID() uint8        { return Snowflake(s).PID() }
func (s RoleID) Increment() uint16 { return Snowflake(s).Increment() }

// StageID is the snowflake type for a StageID.
type StageID Snowflake

// NullStageID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullStageID = StageID(NullSnowflake)

func (s StageID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *StageID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s StageID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s StageID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s StageID) IsNull() bool { return Snowflake(s).IsNull() }

func (s StageID) Time() time.Time   { return Snowflake(s).Time() }
func (s StageID) Worker() uint8     { return Snowflake(s).Worker() }
func (s StageID) PID() uint8        { return Snowflake(s).PID() }
func (s StageID) Increment() uint16 { return Snowflake(s).Increment() }

// StickerID is the snowflake type for a StickerID.
type StickerID Snowflake

// NullStickerID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullStickerID = StickerID(NullSnowflake)

func (s StickerID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *StickerID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s StickerID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s StickerID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s StickerID) IsNull() bool { return Snowflake(s).IsNull() }

func (s StickerID) Time() time.Time   { return Snowflake(s).Time() }
func (s StickerID) Worker() uint8     { return Snowflake(s).Worker() }
func (s StickerID) PID() uint8        { return Snowflake(s).PID() }
func (s StickerID) Increment() uint16 { return Snowflake(s).Increment() }

// StickerPackID is the snowflake type for a StickerPackID.
type StickerPackID Snowflake

// NullStickerPackID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullStickerPackID = StickerPackID(NullSnowflake)

func (s StickerPackID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *StickerPackID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s StickerPackID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s StickerPackID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s StickerPackID) IsNull() bool { return Snowflake(s).IsNull() }

func (s StickerPackID) Time() time.Time   { return Snowflake(s).Time() }
func (s StickerPackID) Worker() uint8     { return Snowflake(s).Worker() }
func (s StickerPackID) PID() uint8        { return Snowflake(s).PID() }
func (s StickerPackID) Increment() uint16 { return Snowflake(s).Increment() }

// TagID is the snowflake type for a TagID.
type TagID Snowflake

// NullTagID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullTagID = TagID(NullSnowflake)

func (s TagID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *TagID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s TagID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s TagID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s TagID) IsNull() bool { return Snowflake(s).IsNull() }

func (s TagID) Time() time.Time   { return Snowflake(s).Time() }
func (s TagID) Worker() uint8     { return Snowflake(s).Worker() }
func (s TagID) PID() uint8        { return Snowflake(s).PID() }
func (s TagID) Increment() uint16 { return Snowflake(s).Increment() }

// TeamID is the snowflake type for a TeamID.
type TeamID Snowflake

// NullTeamID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullTeamID = TeamID(NullSnowflake)

func (s TeamID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *TeamID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s TeamID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s TeamID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s TeamID) IsNull() bool { return Snowflake(s).IsNull() }

func (s TeamID) Time() time.Time   { return Snowflake(s).Time() }
func (s TeamID) Worker() uint8     { return Snowflake(s).Worker() }
func (s TeamID) PID() uint8        { return Snowflake(s).PID() }
func (s TeamID) Increment() uint16 { return Snowflake(s).Increment() }

// UserID is the snowflake type for a UserID.
type UserID Snowflake

// NullUserID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullUserID = UserID(NullSnowflake)

func (s UserID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *UserID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s UserID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s UserID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s UserID) IsNull() bool { return Snowflake(s).IsNull() }

func (s UserID) Time() time.Time   { return Snowflake(s).Time() }
func (s UserID) Worker() uint8     { return Snowflake(s).Worker() }
func (s UserID) PID() uint8        { return Snowflake(s).PID() }
func (s UserID) Increment() uint16 { return Snowflake(s).Increment() }

// WebhookID is the snowflake type for a WebhookID.
type WebhookID Snowflake

// NullWebhookID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullWebhookID = WebhookID(NullSnowflake)

func (s WebhookID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *WebhookID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s WebhookID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s WebhookID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s WebhookID) IsNull() bool { return Snowflake(s).IsNull() }

func (s WebhookID) Time() time.Time   { return Snowflake(s).Time() }
func (s WebhookID) Worker() uint8     { return Snowflake(s).Worker() }
func (s WebhookID) PID() uint8        { return Snowflake(s).PID() }
func (s WebhookID) Increment() uint16 { return Snowflake(s).Increment() }

// EventID is the snowflake type for a EventID.
type EventID Snowflake

// NullEventID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullEventID = EventID(NullSnowflake)

func (s EventID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *EventID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s EventID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s EventID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s EventID) IsNull() bool { return Snowflake(s).IsNull() }

func (s EventID) Time() time.Time   { return Snowflake(s).Time() }
func (s EventID) Worker() uint8     { return Snowflake(s).Worker() }
func (s EventID) PID() uint8        { return Snowflake(s).PID() }
func (s EventID) Increment() uint16 { return Snowflake(s).Increment() }

// EntityID is the snowflake type for a EntityID.
type EntityID Snowflake

// NullEntityID gets encoded into a null. This is used for optional and nullable snowflake fields.
const NullEntityID = EntityID(NullSnowflake)

func (s EntityID) MarshalJSON() ([]byte, error)  { return Snowflake(s).MarshalJSON() }
func (s *EntityID) UnmarshalJSON(v []byte) error { return (*Snowflake)(s).UnmarshalJSON(v) }

// String returns the ID, or nothing if the snowflake isn't valid.
func (s EntityID) String() string { return Snowflake(s).String() }

// IsValid returns whether or not the snowflake is valid.
func (s EntityID) IsValid() bool { return Snowflake(s).IsValid() }

// IsNull returns whether or not the snowflake is null. This method is rarely
// ever useful; most people should use IsValid instead.
func (s EntityID) IsNull() bool { return Snowflake(s).IsNull() }

func (s EntityID) Time() time.Time   { return Snowflake(s).Time() }
func (s EntityID) Worker() uint8     { return Snowflake(s).Worker() }
func (s EntityID) PID() uint8        { return Snowflake(s).PID() }
func (s EntityID) Increment() uint16 { return Snowflake(s).Increment() }