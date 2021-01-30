# connect.py
# derived from this tutorial: https://realpython.com/how-to-make-a-discord-bot-python/#how-to-make-a-discord-bot-in-the-developer-portal
import os
import sys
import random
from discord.ext import commands        # pip install -U discord.py
from dotenv import load_dotenv          # pip install -U python-dotenv

# Load our discord tokens from the .env file
load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')
GUILD = os.getenv('DISCORD_GUILD')
CHANNEL = os.getenv('DISCORD_CHANNEL')

# Create a bot to interact with Discord
bot = commands.Bot(command_prefix='!')

# Make an event handler for when the bot is ready
@bot.event
async def on_ready():
    guild = bot.guilds[0]                # This is not good code ! lol
    print(
        f'{bot.user.name} is connected to the following guild:\n'
        f'{guild.name}(id: {guild.id})\n'
    )

    # for-loops in Python are a little different; they use 'in' and
    # 'range'. So, for each member in the array guild.members,
    # print the member range. If you want a range thats from 0 - 6 (for ex),
    # that'd look a little like this: for x in range(6):
    members = '\n - '.join([member.name for member in guild.members])
    print(f'Guild Members:\n - {members}')

    channel = guild.text_channels[0]        # This is not good code ! lol
    await channel.send('WHATS UP BITCHES IM A BOT')

# Create a command for the bot to respond to
@bot.command(name='bot')
async def send_message(ctx):
    responses = [
        'Hi !! im a bot. How are you?',
        'I cant provide any meaningful responses currently'
    ]

    response = random.choice(responses)
    await message.channel.send(response)

# Create a command for a dice roll
@bot.command(name='roll_dice', help='Simulates rolling dice.')
async def roll(ctx, number_of_dice: int, number_of_sides: int):
    dice = [
        str(random.choice(range(1, number_of_sides + 1)))
        for _ in range(number_of_dice)
    ]
    await ctx.send(', '.join(dice))

# Give your bot the role 'admin' from the discord app
# Create a command listener for creating a channel
@bot.command(name='create-channel')
@commands.has_role('admin')
async def create_channel(ctx, channel_name='the-bots-channel'):
    guild = ctx.guild
    existing_channel = discord.utils.get(guild.channels, name=channel_name)
    if not existing_channel:
        print(f'Creating a new channel: {channel_name}')
        await guild.create_text_channel(channel_name)

bot.run(TOKEN)

# Add an error handler
@bot.event
async def on_error(event, *args, **kwargs):
    with open('err.log', 'a') as f:
        if event == 'on_message':
            f.write(f'Unhandled message: {args[0]}\n')
        else:
            raise

bot.run(TOKEN)
