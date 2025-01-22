# Unexpected Zero Value When Accessing Non-Existent Map Keys in Go

This example demonstrates a common, yet subtle, issue in Go when working with maps: accessing a non-existent key returns the zero value for the map's value type. This can lead to unexpected behavior if not carefully considered.

## The Problem

Go maps, unlike some other languages, do not throw an error when you attempt to access a key that doesn't exist. Instead, they simply return the zero value of the map's value type. This means that if your map's value type is `int`, accessing a non-existent key will return `0`. If it's a `string`, it will return an empty string, and so on.

This behavior can be problematic if you're not checking for the existence of the key before accessing it, because you might end up using a default zero value instead of realizing that the key is actually missing.

## Solution

The safest way to avoid this issue is to always check if a key exists before accessing it using the `value, ok := map[key]` idiom.  The `ok` boolean variable will be `true` if the key exists, and `false` otherwise. This allows for explicit handling of the case where the key is not found.
