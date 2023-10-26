// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createApplication = `-- name: CreateApplication :exec
INSERT INTO applications (
    appid,
    appeui,
    description,
    devices
) VALUES (
    $1, $2, $3, $4
)
`

type CreateApplicationParams struct {
	Appid       int32
	Appeui      string
	Description sql.NullString
	Devices     sql.NullInt32
}

func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) error {
	_, err := q.db.ExecContext(ctx, createApplication,
		arg.Appid,
		arg.Appeui,
		arg.Description,
		arg.Devices,
	)
	return err
}

const createGateway = `-- name: CreateGateway :exec
INSERT INTO gateways (
    gwid,
    gweui,
    description,
    last_seen
) VALUES (
    $1, $2, $3, $4
)
`

type CreateGatewayParams struct {
	Gwid        int32
	Gweui       string
	Description sql.NullString
	LastSeen    sql.NullTime
}

func (q *Queries) CreateGateway(ctx context.Context, arg CreateGatewayParams) error {
	_, err := q.db.ExecContext(ctx, createGateway,
		arg.Gwid,
		arg.Gweui,
		arg.Description,
		arg.LastSeen,
	)
	return err
}

const deleteApplication = `-- name: DeleteApplication :exec
DELETE FROM applications
WHERE appid = $1
`

func (q *Queries) DeleteApplication(ctx context.Context, appid int32) error {
	_, err := q.db.ExecContext(ctx, deleteApplication, appid)
	return err
}

const getApplicationByID = `-- name: GetApplicationByID :one
SELECT appid, appeui, devices, description FROM applications
WHERE appid = $1
LIMIT 1
`

func (q *Queries) GetApplicationByID(ctx context.Context, appid int32) (Application, error) {
	row := q.db.QueryRowContext(ctx, getApplicationByID, appid)
	var i Application
	err := row.Scan(
		&i.Appid,
		&i.Appeui,
		&i.Devices,
		&i.Description,
	)
	return i, err
}

const getNetworkSettings = `-- name: GetNetworkSettings :one
SELECT id, network_type, beacon_timing, beacon_period, beacon_slot_length, beacon_channels, beacon_dr, channel_settings FROM network_settings
WHERE network_type = $1
LIMIT 1
`

func (q *Queries) GetNetworkSettings(ctx context.Context, networkType string) (NetworkSetting, error) {
	row := q.db.QueryRowContext(ctx, getNetworkSettings, networkType)
	var i NetworkSetting
	err := row.Scan(
		&i.ID,
		&i.NetworkType,
		&i.BeaconTiming,
		&i.BeaconPeriod,
		&i.BeaconSlotLength,
		&i.BeaconChannels,
		&i.BeaconDr,
		&i.ChannelSettings,
	)
	return i, err
}

const insertOrUpdateNetworkSettings = `-- name: InsertOrUpdateNetworkSettings :exec
INSERT INTO network_settings (
    network_type,
    beacon_timing,
    beacon_period,
    beacon_slot_length,
    beacon_channels,
    beacon_dr,
    channel_settings
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (network_type)
DO UPDATE
SET beacon_timing = EXCLUDED.beacon_timing,
    beacon_period = EXCLUDED.beacon_period,
    beacon_slot_length = EXCLUDED.beacon_slot_length,
    beacon_channels = EXCLUDED.beacon_channels,
    beacon_dr = EXCLUDED.beacon_dr,
    channel_settings = EXCLUDED.channel_settings
`

type InsertOrUpdateNetworkSettingsParams struct {
	NetworkType      string
	BeaconTiming     int32
	BeaconPeriod     string
	BeaconSlotLength string
	BeaconChannels   string
	BeaconDr         int32
	ChannelSettings  json.RawMessage
}

func (q *Queries) InsertOrUpdateNetworkSettings(ctx context.Context, arg InsertOrUpdateNetworkSettingsParams) error {
	_, err := q.db.ExecContext(ctx, insertOrUpdateNetworkSettings,
		arg.NetworkType,
		arg.BeaconTiming,
		arg.BeaconPeriod,
		arg.BeaconSlotLength,
		arg.BeaconChannels,
		arg.BeaconDr,
		arg.ChannelSettings,
	)
	return err
}

const listApplications = `-- name: ListApplications :many
SELECT appid, appeui, devices, description FROM applications
`

func (q *Queries) ListApplications(ctx context.Context) ([]Application, error) {
	rows, err := q.db.QueryContext(ctx, listApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Application
	for rows.Next() {
		var i Application
		if err := rows.Scan(
			&i.Appid,
			&i.Appeui,
			&i.Devices,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGateway = `-- name: ListGateway :many
SELECT gwid, gweui, description, last_seen FROM gateways
`

func (q *Queries) ListGateway(ctx context.Context) ([]Gateway, error) {
	rows, err := q.db.QueryContext(ctx, listGateway)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Gateway
	for rows.Next() {
		var i Gateway
		if err := rows.Scan(
			&i.Gwid,
			&i.Gweui,
			&i.Description,
			&i.LastSeen,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
