<script>
    export let owner = false;
</script>

<div class="flex gap-5 mb-5 {owner ? 'flex-row-reverse' : ''}">
    <div class="flex flex-col font-light {owner ? 'items-end' : ''}">
        <img
            src="https://robohash.org/ab6ef9bb0974d105efd6d0cab1eec20a?set=set2&bgset=bg2&size=200x200"
            alt=""
            class="w-10 h-10 rounded-full object-cover"
        />
        <span>Just now{owner}</span>
    </div>
    <div class="max-w-[50%] flex flex-col gap-2 {owner ? 'items-end' : ''}">
        <p
            class="py-2 px-5 bg-white rounded-lg max-w-max {owner
                ? 'rounded-tr-none'
                : 'rounded-tl-none'}"
        >
            helosssssssssssssssssssss
        </p>

        <img
            src="https://robohash.org/ab6ef9bb0974d105efd6d0cab1eec20a?set=set2&bgset=bg2&size=200x200"
            alt=""
            class="w-2/4"
        />
    </div>
</div>
